import gridView from './grid-view'

const hasGrid = (): boolean => {
    const cardGridInners = document.getElementsByClassName('card-grid-inner')
    return cardGridInners.length > 0
}

export default () => {
    if (hasGrid()) {
        gridView()
        window.addEventListener("resize", gridView);
    }
}