export interface EventTypeList {
    base: Base;
    data: Data;
    [property: string]: any;
}

export interface Base {
    code: number;
    message: string;
    [property: string]: any;
}

export interface Data {
    list: EventType[];
    [property: string]: any;
}

export interface EventType {
    types: EventTypeItem[];
    name: string;
    [property: string]: any;
}

export interface EventTypeItem {
    id: string;
    name: string;
    [property: string]: any;
}
