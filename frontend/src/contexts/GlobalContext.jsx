
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

    // ====== Selected Stack state
    const [showSelectedItemsStack, setShowSelectedItemsStack] = useState(false);
    
    const toggleSelectedItemsStack = () => {
        setShowSelectedItemsStack(!showSelectedItemsStack);
    }

    const [selectedItemsStack, setSelectedItemsStack] = useState([]);
    
    const pushToSelectedStack = (item) => {
        setSelectedItemsStack([...selectedItemsStack, ...item]);
    }


    const value = {
        showSidebar,
        toggleSidebar,

        sliderValue,
        setSliderValue,

        showSelectedItemsStack,
        toggleSelectedItemsStack,

        selectedItemsStack,
        setSelectedItemsStack,
        pushToSelectedStack,
    }

    return (
        <GlobalContext.Provider value={value}>
            {children}
        </GlobalContext.Provider>
    )
}

export { GlobalContext, GlobalProvider }