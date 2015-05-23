import 'dart:html';

const apiBaseUrl = "http://localhost:3000/api";

main() async{
  var res = await HttpRequest.getString(apiBaseUrl);
  querySelector('#output').text = res;
}
