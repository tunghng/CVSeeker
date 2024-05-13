
import { useState, createContext } from 'react';

const GlobalContext = createContext();

function GlobalProvider({ children }) {

    // ====== Left Sidebar state
    const [showSidebar, setShowSidebar] = useState(true);
    
    const toggleSidebar = () => {
        setShowSidebar(!showSidebar);
    }


    // ====== Selected Stack bar state
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
    
    // ====== Detailed Item Modal state
    const [showDetailItemModal, setShowDetailItemModal] = useState(false);
    const [detailItem, setDetailItem] = useState(null);


    const value = {
        showSidebar,
        toggleSidebar,

        showSelectedItemsStack,
        toggleSelectedItemsStack,

        selectedItemsStack,
        setSelectedItemsStack,
        pushToSelectedStack,
        popFromSelectedStack,

        showDetailItemModal,
        setShowDetailItemModal,
        detailItem,
        setDetailItem,
    }

    return (
        <GlobalContext.Provider value={value}>
            {children}
        </GlobalContext.Provider>
    )
}

export { GlobalContext, GlobalProvider }