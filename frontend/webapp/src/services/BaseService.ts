export default class BaseService
{
    joinUrl(...args: string[]): string
    {
        return args.join("/");
    }
}