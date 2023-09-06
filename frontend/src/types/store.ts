
export interface ISpell {
    id: number,
    name: string,
    description: string,
    summonerLevel: number,
    cooldown: number,
    gameModes: string[],
    iconPath: string,
}

export interface IQueue {
    name: string,
    shortName: string,
    description: string,
    id: number,
    detailedDescription: string
}

export interface IChampion {
    id: number;
    name: string;
    alias: string;
    squarePortraitPath: string;
    roles: string[];
}

export interface IItem {
    id: number,
    name: string,
    description: string,
    active: boolean,
    inStore: boolean,
    from: number[],
    to: number[],
    categories: string[],
    maxStacks: number,
    requiredChampion: string,
    requiredAlly: string,
    requiredBuffCurrencyName: string,
    requiredBuffCurrencyCost: number,
    specialRecipe: number,
    isEnchantment: boolean,
    price: number,
    priceTotal: number,
    iconPath: string,
}

export interface IPerk {
    id: number,
    name: string,
    majorChangePatchVersion: string,
    tooltip: string,
    shortDesc: string,
    longDesc: string,
    recommendationDescriptor: string,
    iconPath: string,
    endOfGameStatDescs: string[],
    recommendationDescriptorAttributes: {},
}