import sysinfo from 'systeminformation'
import isDocker from 'is-docker'
import isWSL from 'is-wsl'

export default class Machine {
  public static cpu = async () => ({
    ...(await sysinfo.cpu()),
    cpuTemp: await sysinfo.cpuTemperature(),
    curentSpeed: await sysinfo.cpuCurrentSpeed(),
  })

  public static battery = () => sysinfo.battery()

  public static time = () => sysinfo.time()

  public static isDocker = () => isDocker()

  public static isRoot = () => process.getuid && process.getuid() === 0

  /**
   * Detects Windows Subsystem for Linux.
   */
  public static isWSL = () => isWSL
}