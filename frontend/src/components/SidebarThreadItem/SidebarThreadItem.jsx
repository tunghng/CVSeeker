
import { useState, useEffect, useRef } from "react";
import { Link } from "react-router-dom";
import FeatherIcon from "feather-icons-react";

const SidebarThreadItem = ({ item, isActive }) => {
    const [showPopup, setShowPopup] = useState(false);
    const popupRef = useRef(null);

    const moreButtonClickHandler = (e) => {
        e.preventDefault();
        e.stopPropagation();
        setShowPopup(true);
    };

    const renameHandler = () => {
        setShowPopup(false);
        // Implement rename functionality
    };

    const deleteHandler = () => {
        setShowPopup(false);
        // Implement delete functionality
    };

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (popupRef.current && !popupRef.current.contains(event.target)) {
                setShowPopup(false);
            }
        };
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);

    return (
        <Link to={`/chat/${item.id}`} className={`thread-item relative group ${(isActive || showPopup) && 'active'}`}>
            <span>{item.name === '' ? 'New Thread' : item.name}</span>

            <button
                className={`ml-3 rounded-md hidden text-text group-hover:block ${showPopup && '!block'} hover:opacity-80 transition-all duration-300 ease-in-out`}
                onClick={moreButtonClickHandler}
            >
                <FeatherIcon icon="more-horizontal" className="w-8 h-8 p-1.5" strokeWidth={1.8} />
            </button>

            {showPopup && (
                <div ref={popupRef} className="absolute top-5 left-[calc(100%-20px)] bg-white shadow-md rounded-md border border-border p-2 z-10">
                    <MorePopup renameHandler={renameHandler} deleteHandler={deleteHandler} />
                </div>
            )}
        </Link>
    )
}

const MorePopup = ({ renameHandler, deleteHandler }) => {
    return (
        <>
            <button onClick={renameHandler} className="px-3 py-2 w-full flex items-center rounded-md text-sm text-gray-700 hover:bg-gray-100">
                <FeatherIcon icon="edit" className="w-4 h-4 mr-1.5" strokeWidth={1.8} />
                Rename
            </button>
            <button onClick={deleteHandler} className="px-3 py-2 w-full flex items-center rounded-md text-sm text-gray-700 hover:bg-gray-100">
                <FeatherIcon icon="trash-2" className="w-4 h-4 mr-1.5" strokeWidth={1.8} />
                Delete
            </button>
        </>
    );
};

export default SidebarThreadItem