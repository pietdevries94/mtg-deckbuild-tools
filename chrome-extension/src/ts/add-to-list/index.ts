import Vue from 'vue'
import Base from './components/Base.vue'

export interface Card {
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
                currentCardIDs: <Card>{}
            },
        }
    ),
}).$mount(vueDiv)

export const addCardByCodeAndNumber = (set: string, number: string) => {
    Vue.set(vueInstance.$children[0]!.$props, 'currentCardIDs', { set, number })
}
