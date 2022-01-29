package app.services;

import sys.FileSystem;

class Walker {
  public function new () {}

  public static function go () {
    var currentPath = FileSystem.absolutePath('.');

    trace(currentPath);
  }
}