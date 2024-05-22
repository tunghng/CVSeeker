
import MyInput from '../common/MyInput'

const LinkedinUploadInput = ({ value, onChange, onPressEnter, onClickButton }) => {

    const checkValidUrl = (url) => {
        const linkedinRegex = /^(https?:\/\/)?(www\.)?linkedin\.com\/in\/[a-zA-Z0-9_-]+\/?$/;
        return linkedinRegex.test(url);
    }

    return (
        <MyInput
            value={value}
            onChange={onChange}
            onPressEnter={onPressEnter}
            onClickButton={onClickButton}
            placeholder="https://www.linkedin.com/in/account"
            icon="upload"
            text='Upload'
            className='mt-2'
            inputClassName={`${checkValidUrl(value) ? 'border-primary' : '!border-border focus:!border-disable-dark'}`}
            buttonClassName={`${checkValidUrl(value) ? 'bg-primary peer-focus:text-white text-white hover:bg-primary-hover' : '!bg-border !text-gray-500 !cursor-default peer-focus:!bg-disable-dark peer-focus:!text-gray-600'}
                            h-full pl-3 pr-4 rounded-l-none !right-0`}
            iconClassName='!w-5 !h-5'
            disabled={!checkValidUrl(value)}
            autoFocus={false}
        />
    )
}

export default LinkedinUploadInput
