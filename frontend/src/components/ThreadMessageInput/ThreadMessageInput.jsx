
import { useRef, useEffect } from "react";
import FeatherIcon from "feather-icons-react";
import './ThreadMessageInput.css';

const ThreadMessageInput = ({ value, onChange, onPressEnter, onClickButton }) => {
    const textareaRef = useRef(null);

    useEffect(() => {
        const textarea = textareaRef.current;
        let maxHeight = 140;
        if (textarea) {
            textarea.style.height = 'auto';
            maxHeight = textarea.scrollHeight > maxHeight ? maxHeight : textarea.scrollHeight;
            textarea.style.overflowY = textarea.scrollHeight > maxHeight ? 'auto' : 'hidden';
            textarea.style.height = `${maxHeight}px`;
        }
    }, [value]);

    return (
        <div className="thread-message-input-wrapper relative flex items-center w-full my-5 rounded-lg border-2 border-border transition-all duration-300 ease-in-out">
            <textarea
                ref={textareaRef}
                id="thread-message-input"
                className="flex-1 px-3 py-1 my-2 mr-14 peer bg-transparent text-text font-medium text-base outline-none 
                           resize-none"
                placeholder="Type a message..."
                value={value}
                onChange={onChange}
                onKeyDown={onPressEnter}
                rows={1}
            />

            <button
                onClick={onClickButton}
                className="absolute right-3 bottom-2 p-1.5 my-button my-button-subtle flex items-center rounded-md transition-all duration-300 ease-in-out"
            >
                <FeatherIcon icon="send" className="w-5 h-5" />
            </button>
        </div>
    );
};

export default ThreadMessageInput;


