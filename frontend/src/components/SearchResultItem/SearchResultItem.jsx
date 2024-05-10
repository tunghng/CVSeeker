
import { useContext } from "react"
import { GlobalContext } from "../../contexts/GlobalContext"

import FeatherIcon from 'feather-icons-react'

const SearchResultItem = ({ item, viewMode, handleItemClick }) => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);

    // ====== Event Handlers ======
    const handleDetailDoubleClick = () => {
        // console.log(item.id);
    }
    const handleDetailClick = (e) => {
        e.stopPropagation();
        globalContext.setSelectedItem(item);
        globalContext.setShowDetailItemModal(true);
    }
    const handleDownloadClick = (e) => {
        e.stopPropagation();
        console.log(item.name);
    }

    return (
        <div className="h-14 flex border-b border-border hover:bg-primary-subtle"
            onDoubleClick={handleDetailDoubleClick}
            onClick={() => handleItemClick(item.id)}
        >
            <div className="w-10 flex justify-center items-center relative">
                <input
                    type="checkbox"
                    checked={item.selected}
                    readOnly
                    className="appearance-none peer w-6 h-6 border-2 border-border rounded-full
                             checked:bg-primary checked:border-primary"
                />
                <FeatherIcon icon="check" className='hidden peer-checked:block absolute top-1/2 -translate-x-1/2 left-1/2 -translate-y-1/2 w-[18px] h-[18px] text-white pointer-events-none' strokeWidth={3} />
            </div>

            <p className='flex-1 flex items-center text-lg font-semibold cursor-default truncate'>{item.name}</p>

            <p className='w-32 flex items-center justify-center text-sm cursor-default'>07-07-2024 17:20</p>

            <div className="w-24 sm:w-56 flex items-center justify-center gap-x-1">
                <button
                    className='px-3 py-2 sm:py-1 bg-disable-light text-base rounded-full flex items-center gap-x-1 hover:bg-secondary-subtle'
                    onClick={handleDetailClick}>
                    <FeatherIcon icon="info" className='w-[18px] h-[18px]' strokeWidth={1.6} />
                    <p className='hidden sm:block'>Detail</p>
                </button>
                <button
                    className='px-3 py-2 sm:py-1 bg-disable-light text-base rounded-full flex items-center gap-x-1 hover:bg-secondary-subtle'
                    onClick={handleDownloadClick}>
                    <FeatherIcon icon="download" className='w-[18px] h-[18px]' strokeWidth={1.6} />
                    <p className='hidden sm:block'>Download</p>
                </button>
            </div>
        </div>
    );
}

export default SearchResultItem
