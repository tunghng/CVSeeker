
import { useState, useContext, useEffect } from "react"
import { useParams } from "react-router-dom"
import { GlobalContext } from "../contexts/GlobalContext"
import getThreadMessage from "../services/chat/getThreadMessage"
import { v4 as uuidv4 } from 'uuid';

import StackItem from "../components/StackItem/StackItem"
import DetailItemModal from "../components/DetailItemModal/DetailItemModal"
import FeatherIcon from 'feather-icons-react'
import { Tooltip } from "react-tooltip"
import ThreadMessageList from "../components/ThreadMessageList/ThreadMessageList"
import ThreadMessageInput from "../components/ThreadMessageInput/ThreadMessageInput"

const ChatPage = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    let { threadId } = useParams();
    const [threadInfo, setThreadInfo] = useState(null);
    const [threadMessages, setThreadMessages] = useState([]);
    const [threadInput, setThreadInput] = useState('');

    // ====== Fetching Thread Messages ======
    useEffect(() => {
        setThreadMessages([])
        getThreadMessage(threadId)
            .then(res => {
                setThreadInfo(res)
                setThreadMessages(res.data)
            })
    }, [threadId]);

    // ====== Event Handlers ======
    const stackItemDetailClickHandler = (item) => {
        globalContext.setDetailItem(item)
        globalContext.setShowDetailItemModal(true)
    }
    const stackItemRemoveClickHandler = (itemId) => {
        globalContext.popFromSelectedStack(itemId)
    }
    const detailItemModalCloseHandler = () => {
        globalContext.setShowDetailItemModal(false)
    }

    const threadMessageSendKeyDownHandler = (e) => {
        if (e.key === 'Enter') {
            if (e.shiftKey) {
                setThreadInput(threadInput + '\n');
            }
            else if (threadInput.trim() !== '') {
                const newMessage = {
                    id: uuidv4(),
                    role: 'user',
                    content: [
                        {
                            type: "text",
                            text: {
                                value: threadInput.trim().replace(/\n/g, '\n\n')
                            }
                        }
                    ]
                };
                setThreadMessages([...threadMessages, newMessage]);
                setThreadInput('');
            }
            e.preventDefault();
        }
    };

    const threadMessageSendClickHandler = () => {
        if (threadInput.trim() !== '') {
            const newMessage = {
                id: uuidv4(),
                role: 'user',
                content: [
                    {
                        type: "text",
                        text: {
                            value: threadInput.trim().replace(/\n/g, '\n\n')
                        }
                    }
                ]
            };
            setThreadMessages([...threadMessages, newMessage]);
            setThreadInput('');
        }
    };

    return (
        <main className="my-content-wrapper flex">

            {/* ====== Thread Messages ====== */}
            <div className={`${globalContext.showSelectedItemsStack && 'md:mr-72'} flex-1 transition-all duration-700 ease-in-out`}>
                <div className="my-container-medium mt-6 pb-28">
                    {
                        (threadMessages === null || threadMessages.length === 0) ?
                            (
                                <div className="flex flex-col items-center justify-center h-[calc(100%-3rem)] space-y-4">
                                    <p className="text-text">Loading messsages ...</p>
                                    <div className="loader"></div>
                                </div>
                            ) :
                            <ThreadMessageList threadMessages={threadMessages} />
                    }
                </div>
            </div>


            {/* ====== Thread Input ====== */}
            <div className={`fixed bottom-0 ${globalContext.showSelectedItemsStack ? 'right-72' : 'right-0'} ${globalContext.showSidebar ? 'left-64' : 'left-0'} pr-4 transition-all duration-700 ease-in-out`}>
                <div className="my-container-medium bg-background">
                    <ThreadMessageInput
                        value={threadInput}
                        onChange={(e) => setThreadInput(e.target.value)}
                        onPressEnter={threadMessageSendKeyDownHandler}
                        onClickButton={threadMessageSendClickHandler}
                    />
                </div>
            </div>




            {/* ====== Selected Items Stack ====== */}
            <div className={`${globalContext.showSelectedItemsStack ? 'translate-x-0' : 'translate-x-full'} w-full max-w-72 h-[calc(100%-3rem)] fixed  right-0 flex flex-col bg-background px-3 pt-3 pb-5 border-l-2 border-border transition-all duration-700 ease-in-out`}>
                <h1 className="text-lg font-semibold">Selected items ({globalContext.selectedItemsStack.length})</h1>

                <div className="flex-1">
                    {
                        globalContext.selectedItemsStack.map(item => (
                            <StackItem
                                key={item.id}
                                item={item}
                                onDetailClick={stackItemDetailClickHandler}
                                onRemoveClick={stackItemRemoveClickHandler}
                                showRemoveIcon={false}
                            />
                        ))
                    }
                </div>


                {/* ====== Toggle Stack Button ====== */}
                <button
                    className="absolute top-1/2 -left-6 transform -translate-x-1/2 -translate-y-1/2"
                    onClick={globalContext.toggleSelectedItemsStack}
                >
                    <div className="flex h-12 w-6 flex-col items-center justify-center group"
                        data-tooltip-id="toggle-stack-tooltip"
                        data-tooltip-content={globalContext.showSelectedItemsStack ? 'Close Stackbar' : 'Open Stackbar'}
                        data-tooltip-place="left"
                        data-tooltip-delay-show={200}>
                        <div className={`${globalContext.showSelectedItemsStack ? 'rotate-0 group-hover:rotate-[-24deg]' : 'rotate-[24deg]'} selected-item-stack-button translate-y-[0.15rem]`}></div>
                        <div className={`${globalContext.showSelectedItemsStack ? 'rotate-0 group-hover:rotate-[24deg]' : 'rotate-[-24deg]'} selected-item-stack-button translate-y-[-0.15rem]`}></div>
                    </div>
                    <Tooltip id="toggle-stack-tooltip" />
                </button>

            </div>


            {/* ====== Detail Item Modal ====== */}
            <DetailItemModal
                showDetailItemModal={globalContext.showDetailItemModal}
                detailItem={globalContext.detailItem}
                onModalClose={detailItemModalCloseHandler}
            />
        </main>
    )
}

export default ChatPage
