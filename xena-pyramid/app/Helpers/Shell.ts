import { execSync } from 'child_process'

export const exe = (command: string) =>
	execSync(command).toString('utf-8')