import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';

@Day(7)
class DaySevenRunner implements DayRunner {
  @override
  String run() {
    var inputReader = AdventInputReader("seven.txt");
    var diagram = inputReader.readIntoLines();
    var tracker = TachyonBeamTracker(diagram);
    return "\nSplitter Count: ${tracker.splitCount}";
  }
}

class TachyonBeamTracker {
  final beamGraph = BeamGraph();
  final List<String> beamDiagram;

  TachyonBeamTracker(this.beamDiagram) {
    var startNode = findStartNode();
    startNode.leftNodeId = findSplitterNodes(
      startNode,
      startNode.x,
      startNode.y,
    )?.getId();
  }

  int get splitCount => beamGraph.splitterNodes.length;

  StartNode findStartNode() {
    for (int y = 0; y < beamDiagram.length; y++) {
      for (int x = 0; x < beamDiagram[0].length; x++) {
        var char = beamDiagram[y][x];
        if (char == 'S') {
          var startNode = StartNode(x, y);
          beamGraph.addNode(startNode);
          return startNode;
        }
      }
    }
    throw StateError("Couldn't find start node!");
  }

  BeamNode? findSplitterNodes(BeamNode parentNode, int x, int startY) {
    if (x < 0 || x >= beamDiagram[0].length) return null;

    for (int y = startY; y < beamDiagram.length; y++) {
      var char = beamDiagram[y][x];
      if (char != "^") continue;

      var node = SplitterNode(x, y, parentNodeId: parentNode.getId());

      if (beamGraph.hasNode(node)) {
        return node;
      }

      beamGraph.addNode(node);

      var leftChildNode = findSplitterNodes(node, x - 1, y);
      if (leftChildNode != null) {
        node.leftNodeId = leftChildNode.getId();
      }

      var rightChildNode = findSplitterNodes(node, x + 1, y);
      if (rightChildNode != null) {
        node.rightNodeId = rightChildNode.getId();
      }

      return node;
    }

    return null;
  }
}

class BeamGraph {
  final Map<String, BeamNode> _nodes = {};

  void addNode(BeamNode node) {
    var id = node.getId();
    if (_nodes.containsKey(id)) {
      throw StateError("Node with ID $id already exists in graph!");
    }
    _nodes[id] = node;
  }

  bool hasNode(BeamNode node) => _nodes.containsKey(node.getId());

  Map<String, BeamNode> get nodes => _nodes;

  BeamNode? get startNode => _nodes['start'];

  Map<String, BeamNode> get splitterNodes =>
      Map<String, BeamNode>.from(_nodes)
        ..removeWhere((key, _) => key == "start");

  int get size => _nodes.length;
}

abstract class BeamNode {
  final int x, y;
  String? parentNodeId;
  String? leftNodeId;
  String? rightNodeId;

  BeamNode(
    this.x,
    this.y, {
    this.parentNodeId,
    this.leftNodeId,
    this.rightNodeId,
  });

  bool get hasParent => parentNodeId != null;
  bool get hasLeft => leftNodeId != null;
  bool get hasRight => rightNodeId != null;

  String getId();

  @override
  String toString() {
    return "Node { id: ${getId()}, x: $x, y: $y"
        "${hasParent ? ", parentNodeId: $parentNodeId" : ""}"
        "${hasLeft ? ", leftNodeId: $leftNodeId" : ""}"
        "${hasRight ? ", rightNodeId: $rightNodeId" : ""}"
        " }";
  }
}

class StartNode extends BeamNode {
  StartNode(
    super.x,
    super.y, {
    super.leftNodeId,
    super.rightNodeId,
    super.parentNodeId,
  });

  @override
  String getId() {
    return "start";
  }
}

class SplitterNode extends BeamNode {
  SplitterNode(
    super.x,
    super.y, {
    super.leftNodeId,
    super.rightNodeId,
    super.parentNodeId,
  });

  @override
  String getId() {
    return "x${x}y$y";
  }
}
