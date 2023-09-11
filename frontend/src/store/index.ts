// index.ts

import { createStore, Store } from 'vuex';
import {model} from "../../wailsjs/go/models";
import InitBackendData = model.InitBackendData;


// 创建你的 RootState 类型
export interface RootState {
    currentSummoner: model.Summoner,
    auth: model.Auth,
}

export default createStore<RootState>({
    state: {
        currentSummoner: new model.Summoner(),
        auth:new model.Auth(),
    },
    mutations: {
        SetState(state, initBackendData :InitBackendData) {
            state.auth = initBackendData.auth;
            state.currentSummoner = initBackendData.currentSummoner;
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
