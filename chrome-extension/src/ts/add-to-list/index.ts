import Vue from 'vue'
import Base from './components/Base.vue'
import Toasted from 'vue-toasted'

// @ts-ignore
Vue.use(Toasted)

export interface CardIDs {
    set?: string
    number?: string
}

const vueDiv = document.createElement('div')
document.
    getElementsByTagName('body').
    item(0)!.
    appendChild(vueDiv)

const vueInstance = new Vue({
    render: (h) => h(
        Base,
        {
            props: {
                currentCardIDs: <CardIDs>{}
            },
        }
    ),
}).$mount(vueDiv)

export const addCardByCodeAndNumber = (set: string, number: string) => {
    Vue.set(vueInstance.$children[0]!.$props, 'currentCardIDs', { set, number })
}
