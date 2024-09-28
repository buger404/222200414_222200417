export interface EventTypeList {
    base: Base;
    /**
     * 项目信息
     */
    data: Data;
    [property: string]: any;
}

/**
 * base
 */
export interface Base {
    code: number;
    msg: string;
    [property: string]: any;
}

/**
 * 项目信息
 */
export interface Data {
    types: EventType[];
    [property: string]: any;
}

export interface EventType {
    id: string;
    name: string;
    [property: string]: any;
}