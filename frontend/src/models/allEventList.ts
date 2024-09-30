export interface AllEventList {
    base: Base;
    data: Data;
    [property: string]: any;
}

/**
 * base
 */
export interface Base {
    code: number;
    msg:  string;
    [property: string]: any;
}

export interface Data {
    event: EventGroup[];
    [property: string]: any;
}

export interface EventGroup {
    contests: Contest[];
    group_id: string;
    [property: string]: any;
}

export interface Contest {
    contest_id: string;
    countries:  Country[];
    date:       string;
    status:     number;
    [property: string]: any;
}

/**
 * country
 */
export interface Country {
    flag:  string;
    name:   string;
    rating: string;
    [property: string]: any;
}