import gridView from './grid-view'
import profileView from './profile-view';

const hasGrid = (): boolean => document.getElementsByClassName('card-grid-inner').length > 0

const hasProfile = (): boolean => document.getElementsByClassName('card-profile').length > 0

export default () => {
    if (hasGrid()) {
        gridView()
        window.addEventListener("resize", gridView);
    }

    if (hasProfile()) {
        profileView()
    }
}