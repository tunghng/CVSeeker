
import { useState, createContext } from "react";

const SaveContext = createContext();

function SaveProvider({ children }) {
    const [saveList, setSaveList] = useState(() => {
        let localSaveList = JSON.parse(localStorage.getItem('SavedCVs'));
        return localSaveList || [];
    });

    const isSaved = (cv) => {
        const exists = saveList.some((savedCV) => savedCV.id === cv.id);
        return exists;
    };

    const toggleSaveCV = (cv) => {
        const exists = isSaved(cv);

        if (exists) {
            const updatedList = saveList.filter((savedCV) => savedCV.id !== cv.id);
            setSaveList(updatedList);
            localStorage.setItem('SavedCVs', JSON.stringify(updatedList));
        } else {
            const updatedList = [...saveList, cv];
            setSaveList(updatedList);
            localStorage.setItem('SavedCVs', JSON.stringify(updatedList));
        }
    };

    const value = {
        saveList,
        setSaveList,
        isSaved,
        toggleSaveCV
    };

    return (
        <SaveContext.Provider value={value}>
            {children}
        </SaveContext.Provider>
    );
}

export { SaveContext, SaveProvider };