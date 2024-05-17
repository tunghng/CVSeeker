
import { useState, useContext, useEffect, useRef } from "react"
import { useParams } from "react-router-dom"
import { GlobalContext } from "../contexts/GlobalContext"
import getThreadMessage from "../services/chat/getThreadMessage"
import sendThreadMessage from "../services/chat/sendThreadMessage"
import { v4 as uuidv4 } from 'uuid';

import StackItem from "../components/StackItem/StackItem"
import DetailItemModal from "../components/DetailItemModal/DetailItemModal"
import { Tooltip } from "react-tooltip"
import ThreadMessageList from "../components/ThreadMessageList/ThreadMessageList"
import ThreadMessageInput from "../components/ThreadMessageInput/ThreadMessageInput"
import ReactMarkdown from 'react-markdown'

const ChatPage = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    let { threadId } = useParams();
    const [threadInfo, setThreadInfo] = useState(null);
    const [threadMessages, setThreadMessages] = useState([]);
    const [threadInput, setThreadInput] = useState('');
    const [isAssistantLoading, setIsAssistantLoading] = useState(false);
    const [assistantTempMessage, setAssistantTempMessage] = useState('');

    const messagesEndRef = useRef(null);

    // ====== Fetching Thread Messages ======
    useEffect(() => {
        setThreadMessages([])
        getThreadMessage(threadId)
            .then(res => {
                setThreadInfo(res)
                setThreadMessages(res.data)
            })
    }, [threadId]);

    useEffect(() => {
        if (messagesEndRef.current) {
            messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
        }
    }, [assistantTempMessage, threadMessages]);

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

    const threadMessageSendKeyDownHandler = async (e) => {
        if (e.key === 'Enter') {
            if (e.shiftKey) {
                setThreadInput(threadInput + '\n');
                e.preventDefault();
            } else if (threadInput.trim() !== '') {
                let message = threadInput.trim().replace(/\n/g, '\n\n');
                e.preventDefault();
                setThreadInput('');
                renderUserMessage(message);
                setIsAssistantLoading(true);
                const response = await sendThreadMessage(threadId, message);
                await renderAssistantTempMessage(response);
            }
        }
    };

    const threadMessageSendClickHandler = async () => {
        if (threadInput.trim() !== '') {
            let message = threadInput.trim().replace(/\n/g, '\n\n');
            setThreadInput('');
            renderUserMessage(message);
            setIsAssistantLoading(true);
            const response = await sendThreadMessage(threadId, message);
            await renderAssistantTempMessage(response);
        }
    };

    const renderUserMessage = (message) => {
        const newMessage = {
            id: uuidv4(),
            role: 'user',
            content: [
                {
                    type: "text",
                    text: {
                        value: message
                    }
                }
            ]
        };
        setThreadMessages(threadMessages => [...threadMessages, newMessage]);
    };

    const renderAssistantTempMessage = async (message) => {
        let index = 0;
        let tempMessage = '';
        const intervalId = setInterval(() => {
            if (index < message.length) {
                tempMessage += message[index];
                setAssistantTempMessage(tempMessage);
                index++;
            } else {
                clearInterval(intervalId);
                setIsAssistantLoading(false);
                renderAssistantMessage(message.join(''));
            }
        }, 100);
    };

    const renderAssistantMessage = (message) => {
        const newMessage = {
            id: uuidv4(),
            role: 'assistant',
            content: [
                {
                    type: "text",
                    text: {
                        value: message
                    }
                }
            ]
        };
        setThreadMessages(threadMessages => [...threadMessages, newMessage]);
        setAssistantTempMessage('');
    };


    return (
        <main className="my-content-wrapper flex">

            {/* ====== Thread Messages ====== */}
            <div className={`${globalContext.showSelectedItemsStack && 'md:mr-72'} flex-1 transition-all duration-700 ease-in-out`}>
                <div className="my-container-medium mt-0 pb-24">
                    {(threadMessages === null || threadMessages.length === 0) ? (
                        <div className="mt-6 flex flex-col items-center space-y-4">
                            <p className="text-subtitle">Loading messages ...</p>
                            <div className="loader"></div>
                        </div>
                    ) : (
                        <div className="flex flex-col">
                            <ThreadMessageList threadMessages={threadMessages} />
                            <div ref={messagesEndRef} />
                        </div>
                    )}
                    {isAssistantLoading && (
                        <div className="mt-5 px-4 py-2.5 rounded-xl bg-border w-fit">
                            {assistantTempMessage === '' ? (
                                <div className="loader my-1"></div>
                            ) : (
                                <ReactMarkdown>{assistantTempMessage}</ReactMarkdown>
                            )}
                        </div>
                    )}
                </div>
            </div>


            {/* ====== Thread Input ====== */}
            <div className={`fixed bottom-0 ${globalContext.showSelectedItemsStack ? 'right-72' : 'right-0'} ${globalContext.showSidebar ? 'left-64' : 'left-0'} pr-4 transition-all duration-700 ease-in-out`}>
                <div className="my-container-medium bg-background pb-5 pt-2">
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
