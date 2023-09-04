
export interface IRankDetails {
    division: string;
    highestDivision: string;
    highestTier: string;
    isProvisional: boolean;
    leaguePoints: number;
    losses: number;
    miniSeriesProgress: string;
    previousSeasonAchievedDivision: string;
    previousSeasonAchievedTier: string;
    previousSeasonEndDivision: string;
    previousSeasonEndTier: string;
    provisionalGameThreshold: number;
    provisionalGamesRemaining: number;
    queueType: string;
    ratedRating: number;
    ratedTier: string;
    tier: string;
    warnings: null | Record<string, any>; // 使用 Record<string, any> 表示可能是任何对象
    wins: number;
}



export interface ISummoner {
    accountId: number;
    displayName: string;
    gameName: string;
    internalName: string;
    nameChangeFlag: boolean;
    percentCompleteForNextLevel: number;
    privacy: string;
    profileIconId: number;
    puuid: string;
    rerollPoints: {
        currentPoints: number;
        maxRolls: number;
        numberOfRolls: number;
        pointsCostToRoll: number;
        pointsToReroll: number;
    };
    summonerId: number;
    summonerLevel: number;
    tagLine: string;
    unnamed: boolean;
    xpSinceLastLevel: number;
    xpUntilNextLevel: number;
}

export interface ISummonerDetails {

}