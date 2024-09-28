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
    list: Type[];
    [property: string]: any;
}

export interface EventType {
    types: Item[];
    name: string;
    [property: string]: any;
}

export interface EventTypeItem {
    id: string;
    name: string;
    [property: string]: any;
}
