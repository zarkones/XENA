import app.services.Perm;

class Env {
  public function new () {}

  // Remote address of the Xena-Pyramid, the bot builder.
  public static function xenaPyramidHost (): String {
    return Perm.decode(Perm.encode('http://127.0.0.1:60667'));
  }

  // ID of the software we want to build.
  public static function buildProfileId (): String {
    return Perm.decode(Perm.encode('52c93c88-c5b7-435f-8212-80c6af7c3ad8'));
  }
}