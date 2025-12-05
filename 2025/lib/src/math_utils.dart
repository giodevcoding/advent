import 'dart:math';

double logBase(num x, num base) => log(x) / log(base);

double log10(num x) => logBase(x, 10);

int countDigits(num x) => x.floor() == 0 ? 1 : log10(x.abs()).floor() + 1;

List<int> divisors(int x) {
  x = x.abs();

  if (x == 0) return List.empty();

  final small = <int>[];
  final large = <int>[];

  for (var i = 1; i * i <= x; i++) {
    if (x % i == 0) {
      small.add(i);
      final other = x ~/ i;
      if (other != i) large.add(other);
    }
  }

  return [...small, ...large.reversed];
}
