import * as Helper from 'App/Helpers'

import Url from 'url-parse'

class Git {
  public readonly pathPrefix = '../xena-pyramid-software-build-'

  public clone = (maybeUrl: string, buildId: string) => {
    const url = new Url(maybeUrl)

    const buildPath = this.pathPrefix + buildId

    // Clean the build folder if needed.
    try {
      Helper.Shell.exe(`rm -r ${buildPath}`)
    } catch (e) {}

    const repoCloningOutput = Helper.Shell.exe(`git clone ${url.protocol}//${url.hostname}${url.pathname} ${buildPath}`)

    if (repoCloningOutput.endsWith('not an empty directory.'))
      return 'ALREADY_CLONED'
    else
      return 'CLONED'    
  }
}

export default new Git()