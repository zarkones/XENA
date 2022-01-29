package app.services;

import haxe.macro.Expr;

// Permutation.
class Perm {
  public function new () {}

  public static macro function encode (data: String): Expr {
    var output: String = '';

    for (index in 0...data.length) {
      output += String.fromCharCode(data.charCodeAt(index) + data.length);
    }
    
    return macro $v{output};
  }

  public static function decode (data: String): String {
    var output: String = '';

    for (index in 0...data.length) {
      output += String.fromCharCode(data.charCodeAt(index) - data.length);
    }

    return output;
  }
}