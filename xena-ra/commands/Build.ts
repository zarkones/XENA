import { BaseCommand } from '@adonisjs/core/build/standalone'

export default class Build extends BaseCommand {

  /**
   * Command name is used to run the command
   */
  public static commandName = 'xena:build'

  /**
   * Command description is displayed in the "help" output
   */
  public static description = ''

  public static settings = {
    /**
     * Set the following value to true, if you want to load the application
     * before running the command
     */
    loadApp: false,

    /**
     * Set the following value to true, if you want this command to keep running until
     * you manually decide to exit the process
     */
    stayAlive: false,
  }

  public async run () {
    this.logger.info('Not implemented.')
  }
}
