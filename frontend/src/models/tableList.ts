export interface TableList {
    base: Base;
    /**
     * 对阵数据，必须按照这个方式排序：比赛阶段较早的 在前，相同的按实际显示出来的 从上到下 的顺序排列
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
 * 对阵数据，必须按照这个方式排序：比赛阶段较早的 在前，相同的按实际显示出来的 从上到下 的顺序排列
 */
export interface Data {
    tables: Table[];
    [property: string]: any;
}

export interface Table {
    countries: Country[];
    period: string;
    special?: string;
    title: string;
    winner: number;
    [property: string]: any;
}

/**
 * country
 */
export interface Country {
    flag?: string;
    name: string;
    rating: string;
    [property: string]: any;
}