List<String> chunk(String str, int chunkSize) {
  var result = <String>[];
  for (var i = 0; i < str.length; i += chunkSize) {
    result.add(str.substring(i, (i + chunkSize).clamp(0, str.length)));
  }
  return result;
}
