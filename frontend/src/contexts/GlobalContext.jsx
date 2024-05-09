
import { useState, createContext } from 'react';

const GlobalContext = createContext();

function GlobalProvider({ children }) {

    // ====== Left Sidebar state
    const [showSidebar, setShowSidebar] = useState(true);
    
    const toggleSidebar = () => {
        setShowSidebar(!showSidebar);
    }

    // ====== Search Slider value state
    const [sliderValue, setSliderValue] = useState(0.5);




    const value = {
        showSidebar,
        toggleSidebar,

        sliderValue,
        setSliderValue,
    }

    return (
        <GlobalContext.Provider value={value}>
            {children}
        </GlobalContext.Provider>
    )
}

export { GlobalContext, GlobalProvider }