import 'dart:async';

import 'package:advent/day.dart';
import 'package:analyzer/dart/constant/value.dart';
import 'package:analyzer/dart/element/element.dart';
import 'package:build/build.dart';
import 'package:code_builder/code_builder.dart';
import 'package:dart_style/dart_style.dart';
import 'package:source_gen/source_gen.dart';
import 'package:glob/glob.dart';

class DayRegistryBuilder implements Builder {
  static final _dayChecker = const TypeChecker.typeNamed(
    Day,
    inPackage: 'advent',
  );
  static final _dayRunnerChecker = const TypeChecker.typeNamed(
    DayRunner,
    inPackage: 'advent',
  );

  @override
  Map<String, List<String>> get buildExtensions => {
    r'$package$': ['lib/day_registry.g.dart'],
  };

  @override
  FutureOr<void> build(BuildStep buildStep) async {
    final dayRunnerClasses = <int, List<ClassElement>>{};

    // Find all Dart files in lib/, including lib/src/.
    final dartFiles = buildStep.findAssets(Glob('lib/**.dart'));

    await for (final assetId in dartFiles) {
      if (assetId.path.endsWith('.g.dart')) continue;

      if (!await buildStep.resolver.isLibrary(assetId)) continue;

      final lib = await buildStep.resolver.libraryFor(assetId);

      for (final cls in lib.classes) {
        final isRunner = _dayRunnerChecker.isAssignableFrom(cls);
        if (!isRunner) continue;

        final ann = _dayChecker.firstAnnotationOfExact(
          cls,
          throwOnUnresolved: false,
        );

        if (ann == null) continue;

        final dayValue = _readDayValue(ann);
        if (dayValue == null) continue;

        dayRunnerClasses.putIfAbsent(dayValue, () => []).add(cls);
      }
    }

    final output = _generateLibrary(dayRunnerClasses);
    final formatted = DartFormatter(
      languageVersion: DartFormatter.latestLanguageVersion,
    ).format(output);

    final outId = AssetId(buildStep.inputId.package, 'lib/day_registry.g.dart');

    await buildStep.writeAsString(outId, formatted);
  }

  int? _readDayValue(DartObject ann) {
    return ann.getField('day')?.toIntValue();
  }

  String _generateLibrary(Map<int, List<ClassElement>> dayRunnerClasses) {
    final mapValues = <Object?, Object?>{};

    final days = dayRunnerClasses.keys.toList()..sort();
    for (final day in days) {
      final classes = dayRunnerClasses[day]!;
      final instances = classes
          .map((c) => refer(c.name!).newInstance(const []))
          .toList();

      mapValues[day] = literalList(instances, refer('DayRunner'));
    }

    final lib = Library((b) {
      var importUris = dayRunnerClasses.values
          .expand((classes) => classes)
          .map((cls) => cls.library.uri.toString())
          .toList();
      importUris.add("package:advent/day.dart");

      for (final uri in importUris) {
        b.directives.add(Directive.import(uri));
      }

      b.body.add(
        Field((fb) {
          fb
            ..name = 'dayRunners'
            ..modifier = FieldModifier.final$
            ..type = refer('Map<int, List<DayRunner>>')
            ..assignment = literalMap(
              mapValues,
              refer('int'),
              refer('List<DayRunner>'),
            ).code;
        }),
      );
    });

    final emitter = DartEmitter.scoped(useNullSafetySyntax: true);
    return '${lib.accept(emitter)}';
  }
}
