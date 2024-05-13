
import FeatherIcon from 'feather-icons-react';

const MyInput = ({
    value,
    onChange,
    onPressEnter,
    onClickSearch,
    placeholder = "Search...",
    className = "",
    inputClassName = "",
    buttonClassName = "",
    iconClassName = "",
    icon = "",
}) => {
    return (
        <div className={`relative flex items-center w-full ${className}`}>
            <input
                type="text"
                className={`flex-1 pl-4 pr-11 py-2 peer bg-transparent rounded-full text-text font-medium text-lg outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out ${inputClassName}`}
                placeholder={placeholder}
                value={value}
                onChange={onChange}
                onKeyDown={onPressEnter}
            />

            {
                icon && <button
                    onClick={onClickSearch}
                    className={`absolute right-2 p-2 rounded-full text-text peer-focus:text-primary transition-all duration-300 ease-in-out ${buttonClassName}`}
                >
                    <FeatherIcon icon={icon} className={`w-6 h-6 ${iconClassName}`} />

                </button>
            }
        </div>
    );
};

export default MyInput;
