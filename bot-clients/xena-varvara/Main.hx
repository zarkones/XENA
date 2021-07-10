import app.services.XenaPyramid;
import Env;

class Main {
  static public function main (): Void {
    // Bot builder.
    final xenaPyramid: XenaPyramid = new XenaPyramid(Env.xenaPyramidHost());

    // Download, save and persist binary files.
    xenaPyramid.grabAndSave(Env.buildProfileId());

    trace('Done.');
  }
}