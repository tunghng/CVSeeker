
import { useState, useEffect, useRef, useContext } from "react";
import { GlobalContext } from "../../contexts/GlobalContext"
import renameThreadName from '../../services/chat/renameThreadName'
import deleteThreadMessage from "../../services/chat/deleteThreadMessage";

import { Link, useNavigate } from "react-router-dom";
import FeatherIcon from "feather-icons-react";
import RenameThreadModal from "../RenameThreadModal/RenameThreadModal";
import DeleteThreadModal from "../DeleteThreadModal/DeleteThreadModal";

const SidebarThreadItem = ({ item, isActive }) => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    const [showMorePopup, setShowMorePopup] = useState(false);
    const morePopupRef = useRef(null);
    const [showRenameModal, setShowRenameModal] = useState(false);
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [newName, setNewName] = useState('');
    const navigate = useNavigate();

    // ====== Event Handlers ======
    const moreButtonClickHandler = (e) => {
        e.preventDefault();
        e.stopPropagation();
        setShowMorePopup(true);
    };

    const renameHandler = () => {
        setShowMorePopup(false);
        setShowRenameModal(true);
    };
    const renameModalCloseHandler = () => {
        setShowRenameModal(false);
    };
    const renameModalRenameHandler = () => {
        renameThreadName(item.id, newName)
            .then((res) => {
                if (res) {
                    globalContext.setSidebarThreads((prev) => {
                        const newThreads = prev.map((thread) => {
                            if (thread.id === item.id) {
                                return { ...thread, name: newName };
                            } else {
                                return thread;
                            }
                        });
                        return newThreads;
                    });
                    setShowRenameModal(false);
                }
            });
    };

    const deleteHandler = () => {
        setShowMorePopup(false);
        setShowDeleteModal(true);
    };
    const deleteModalCloseHandler = () => {
        setShowDeleteModal(false);
    };
    const deleteModalDeleteHandler = () => {
        deleteThreadMessage(item.id)
            .then((res) => {
                globalContext.setSidebarThreads((prev) => {
                    const newThreads = prev.filter((thread) => thread.id !== item.id);
                    return newThreads;
                });
                setShowDeleteModal(false);
                navigate('/');
            });
    };

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (morePopupRef.current && !morePopupRef.current.contains(event.target)) {
                setShowMorePopup(false);
            }
        };
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);

    return (
        <>
            <Link to={`/chat/${item.id}`} className={`thread-item relative group ${(isActive || showMorePopup) && 'active'}`}>
                <span>{item.name === '' ? 'New Thread' : item.name}</span>

                <button
                    className={`ml-3 rounded-md hidden text-text group-hover:block ${showMorePopup && '!block'} hover:opacity-80 transition-all duration-300 ease-in-out`}
                    onClick={moreButtonClickHandler}
                >
                    <FeatherIcon icon="more-horizontal" className="w-8 h-8 p-1.5" strokeWidth={1.8} />
                </button>

                {showMorePopup && (
                    <div ref={morePopupRef} className="absolute top-5 left-[calc(100%-20px)] bg-white shadow-md rounded-md border border-border p-2 z-10">
                        <MorePopup renameHandler={renameHandler} deleteHandler={deleteHandler} />
                    </div>
                )}
            </Link>

            {showRenameModal && (
                <RenameThreadModal
                    value={newName}
                    onChange={(e) => setNewName(e.target.value)}
                    onClose={renameModalCloseHandler}
                    onRename={renameModalRenameHandler}
                />
            )}

            {showDeleteModal && (
                <DeleteThreadModal
                    onClose={deleteModalCloseHandler}
                    onDelete={deleteModalDeleteHandler}
                />
            )}
        </>
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