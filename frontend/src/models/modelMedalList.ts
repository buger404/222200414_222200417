export interface MedalList {
    base: Base;
    data: MedalListData[];
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

export interface MedalListData {
    /**
     * 国旗ID
     */
    flag: string;
    /**
     * 奖牌数，金牌、银牌、铜牌、总数
     */
    medals: Medals;
    /**
     * 国家名
     */
    name: string;
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