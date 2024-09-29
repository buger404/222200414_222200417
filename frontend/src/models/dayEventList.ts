export interface DayEventList {
    base: Base;
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

export interface Data {
    events: DayEvent[];
    [property: string]: any;
}

export interface DayEvent {
    countries: Country[];
    event: string;
    group: string;
    /**
     * ID 编号
     */
    id: string;
    time: string;
    winner: number;
    [property: string]: any;
}

/**
 * country
 */
export interface Country {
    flag: string;
    name: string;
    rating: string;
    [property: string]: any;
}