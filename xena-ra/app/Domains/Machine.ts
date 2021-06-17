import sysinfo from 'systeminformation'
import isDocker from 'is-docker'

export default class Machine {
  public static cpu = async () => {
    const others = await sysinfo.cpu()
    const cpuTemp = await sysinfo.cpuTemperature()
    const curentSpeed = await sysinfo.cpuCurrentSpeed()

    return {
      ...others,
      cpuTemp,
      curentSpeed,
    }
  }

  public static battery = () => sysinfo.battery()

  public static time = () => sysinfo.time()

  public static isDocker = () => isDocker()

  public static isRoot = () => process.getuid && process.getuid() === 0
}