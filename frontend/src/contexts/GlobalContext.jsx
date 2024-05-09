
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
    
    const pushToSelectedStack = (selectItems) => {
        const uniqueItems = selectItems.filter(item => !selectedItemsStack.some(existingItem => existingItem.id === item.id));
        setSelectedItemsStack([...selectedItemsStack, ...uniqueItems]);
    }
    
    const popFromSelectedStack = (itemId) => {
        const newStack = selectedItemsStack.filter(item => item.id !== itemId);
        setSelectedItemsStack(newStack);
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
        popFromSelectedStack,
    }

    return (
        <GlobalContext.Provider value={value}>
            {children}
        </GlobalContext.Provider>
    )
}

export { GlobalContext, GlobalProvider }