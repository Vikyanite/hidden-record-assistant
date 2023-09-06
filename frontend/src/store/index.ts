// index.ts

import { createStore, Store } from 'vuex';
import {IChampion, IItem, IPerk, IQueue, ISpell} from "../types/store";



// 创建你的 RootState 类型
export interface RootState {
    champions: { [key: number]: IChampion},
    perks:{ [key: number]: IPerk },
    items:{ [key: number]: IItem, },
    spells:{ [key: number]: ISpell, },
    queues:{ [key: number]: IQueue, },
    auth: {
        token: string;
        port: string;
    }
}

export default createStore<RootState>({
    state: {
        champions: {
            1:{
                "id":1,
                "name": "黑暗之女",
                "alias": "Annie",
                "squarePortraitPath": "/lol-game-data/assets/v1/champion-icons/1.png",
                "roles": [
                    "mage"
                ]
            }
        }, // 这里定义你要存储的对象
        perks: {
           8369: {
               id: 8369,
                name: "先攻",
                majorChangePatchVersion: "11.23",
                tooltip: "在进入与英雄战斗的@GraceWindow.2@秒内，对一名敌方英雄进行的攻击或技能将提供@GoldProcBonus@金币和<b>先攻</b>效果，持续@Duration@秒，来使你对英雄们造成<truedamage>@DamageAmp*100@%</truedamage>额外<truedamage>伤害</truedamage>，并提供<gold>{{ Item_Melee_Ranged_Split }}</gold>该额外伤害值的<gold>金币</gold>。<br><br>冷却时间：<scaleLevel>@Cooldown@</scaleLevel>秒<br><hr><br>已造成的伤害：@f1@<br>已提供的金币：@f2@",
                shortDesc: "在你率先发起与英雄的战斗时，造成9%额外伤害，持续3秒，并基于该额外伤害提供金币。",
                longDesc: "在进入与英雄战斗的0.25秒内，对一名敌方英雄进行的攻击或技能将提供5金币和<b>先攻</b>效果，持续3秒，来使你对英雄们造成<truedamage>9%</truedamage>额外<truedamage>伤害</truedamage>，并提供<gold>100% (远程英雄为70%)</gold>该额外伤害值的<gold>金币</gold>。<br><br>冷却时间：<scaleLevel>25 ~ 15</scaleLevel>秒",
                recommendationDescriptor: "真实伤害，金币收入",
                iconPath: "/lol-game-data/assets/v1/perk-images/Styles/Inspiration/FirstStrike/FirstStrike.png",
                endOfGameStatDescs: [
                    "已造成的伤害：@eogvar1@",
                    "已提供的金币：@eogvar2@"
                ],
                recommendationDescriptorAttributes: {}
            }
        },
        items:{
            1001: {
                "id": 1001,
                "name": "鞋子",
                "description": "<mainText><stats><attention>25</attention>移动速度</stats></mainText><br>",
                "active": false,
                "inStore": true,
                "from": [],
                "to": [
                    3111,
                    3006,
                    3009,
                    3020,
                    3047,
                    3117,
                    3158
                ],
                "categories": [
                    "Boots"
                ],
                "maxStacks": 1,
                "requiredChampion": "",
                "requiredAlly": "",
                "requiredBuffCurrencyName": "",
                "requiredBuffCurrencyCost": 0,
                "specialRecipe": 0,
                "isEnchantment": false,
                "price": 300,
                "priceTotal": 300,
                "iconPath": "/lol-game-data/assets/ASSETS/Items/Icons2D/1001_Class_T1_BootsofSpeed.png"
            }
        },
        spells: {
            1: {
                "id": 1,
                "name": "净化",
                "description": "移除身上的所有限制效果（压制效果和击飞效果除外）和召唤师技能的减益效果，并且若在接下来的3秒里再次被施加限制效果时，新效果的持续时间会减少65%。",
                "summonerLevel": 9,
                "cooldown": 210,
                "gameModes": [
                    "URF",
                    "CLASSIC",
                    "ARSR",
                    "ARAM",
                    "ULTBOOK",
                    "WIPMODEWIP",
                    "TUTORIAL",
                    "DOOMBOTSTEEMO",
                    "PRACTICETOOL",
                    "FIRSTBLOOD",
                    "NEXUSBLITZ",
                    "PROJECT",
                    "ONEFORALL"
                ],
                "iconPath": "/lol-game-data/assets/DATA/Spells/Icons2D/Summoner_boost.png"
            }
        },
        queues: {
         800: {
             "name": "一般",
             "shortName": "一般",
             "description": "一般",
             "detailedDescription": "",
             "id": 800
         },
        },
        auth: {
            token: "",
            port: ""
        },
    },
    mutations: {
        SetState(state, initBackendData) {
            state.auth = initBackendData.auth;
            state.champions = initBackendData.champions;
            state.items = initBackendData.items;
            state.perks = initBackendData.perks;
            state.spells = initBackendData.spells;
            state.queues = initBackendData.queues;
        }
    },
    getters: {
        LCUAPIPrefix: (state) => () :string => {
            return "/LCUAPI"
        },
        LocalAssetPrefix: (state) => () :string => {
            return "/LOCAL"
        },
    },
});
