export interface MedalList {
    base: Base;
    data: Detail;
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

export interface Detail {
    details : MedalListData[];
}

export interface MedalListData {
    /**
     * 国旗ID
     */
    flag: string;
    /**
     * 奖牌数，金牌、银牌、铜牌、总数
     */
    list: Medals;
    /**
     * 国家名
     */
    countryName: string;
    [property: string]: any;
}

/**
 * 奖牌数，金牌、银牌、铜牌、总数
 */
export interface Medals {
    bronze: number;
    gold: number;
    sliver: number;
    total: number;
    [property: string]: any;
}