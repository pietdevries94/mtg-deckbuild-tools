import Vue from 'vue'
import Base from './components/Base.vue'
import Toasted from 'vue-toasted'

// @ts-ignore
Vue.use(Toasted, {
    duration: 5000
})

export interface CardIDs {
    set: string
    number: string
}

const vueDiv = document.createElement('div')
document.
    getElementsByTagName('body').
    item(0)!.
    appendChild(vueDiv)

const vueInstance = new Vue({
    components: { Base },
    render(createElement) {
        return createElement(
            Base,
            {
                props: {
                    currentCardIDs: this.currentCardIDs,
                },
            }
        )
    },
    data: {
        currentCardIDs: { set: '', number: 0 }
    }
}).$mount(vueDiv)

export const addCardByCodeAndNumber = (set: string, number: string) => {
    Vue.set(vueInstance.$data, 'currentCardIDs', { set, number })
}
