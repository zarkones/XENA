import app.services.Perm;

class Env {
  public function new () {}

  public static function extensions (): Array<String> {
    return [
      'txt',
      'xlsx',
      'xls',
      'xlsm',
      'xlt',
      'xltm',
      'csv',
      'doc',
      'docx',
      'dot',
      'dotx',
      'pdf',
      'pot',
      'potm',
      'ppsx',
      'pptm',
      'pptx',
      'key',
      'rtf',
      'wbk',
      'wks',
      'wpd',
      'wp5',
    ];
  }

  // Remote address of the Xena-Rack.
  public static function xenaRackHost (): String {
    return Perm.decode(Perm.encode('http://127.0.0.1:60667'));
  }
}