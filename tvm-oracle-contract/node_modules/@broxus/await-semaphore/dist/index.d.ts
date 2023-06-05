export declare class Semaphore {
    private tasks;
    count: number;
    constructor(count: number);
    private sched;
    acquire(): Promise<() => void>;
    use<T>(f: () => Promise<T>): Promise<T>;
}
export declare class Mutex extends Semaphore {
    constructor();
}
