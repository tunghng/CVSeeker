import FeatherIcon from 'feather-icons-react';

const MyInput = ({
    value,
    onChange,
    onPressEnter,
    onClickButton,
    placeholder = "Search...",
    className = "",
    inputClassName = "",
    buttonClassName = "",
    iconClassName = "",
    icon = "",
    textClassName = "",
    text = "",
    disabled = false
}) => {

    const handleKeyPress = (e) => {
        if (e.key === 'Enter' && onPressEnter && !disabled) {
            onPressEnter(e);
        }
    };

    const handleClickButton = () => {
        if (onClickButton && !disabled) {
            onClickButton();
        }
    };

    return (
        <div className={`relative flex items-center w-full ${className}`}>
            <input
                type="text"
                className={`flex-1 pl-4 pr-11 py-2 peer bg-transparent rounded-full text-title font-medium text-lg outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out ${inputClassName}`}
                placeholder={placeholder}
                value={value}
                onChange={onChange}
                onKeyDown={handleKeyPress}
            />

            {(icon || text) && (
                <button
                    onClick={handleClickButton}
                    className={`absolute right-2 p-2 flex items-center rounded-full text-text peer-focus:text-primary transition-all duration-300 ease-in-out ${buttonClassName}`}
                >
                    <FeatherIcon icon={icon} className={`w-6 h-6 ${iconClassName}`} />
                    {text && <span className={`ml-2 font-semibold ${textClassName}`}>{text}</span>}
                </button>
            )}
        </div>
    );
};

export default MyInput;
