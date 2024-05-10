
import { useContext } from "react"
import { GlobalContext } from "../../contexts/GlobalContext"
import { useLocation } from "react-router-dom"

import FeatherIcon from "feather-icons-react";

const StackItem = ({ item }) => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    const location = useLocation()

    const handleDetailClick = (e) => {
        e.stopPropagation();
        globalContext.setSelectedItem(item);
        globalContext.setShowDetailItemModal(true);
    }
    const handleRemoveClick = (e) => {
        e.stopPropagation();
        globalContext.popFromSelectedStack(item.id);
    }


    return (
        <div className="px-3 py-2 mt-2 flex items-center justify-between rounded-md bg-disable-light cursor-pointer"
            onClick={handleDetailClick}
        >
            <p>{item.name}</p>
            {
                location.pathname === '/search' &&
                <FeatherIcon icon="x" className="w-6 h-6 p-1 rounded-full cursor-pointer hover:bg-subtitle"
                    onClick={handleRemoveClick} />
            }
        </div>
    )
}

export default StackItem