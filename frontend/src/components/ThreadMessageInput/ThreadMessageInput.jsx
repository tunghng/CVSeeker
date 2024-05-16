
import MyInput from '../common/MyInput'

const ThreadMessageInput = ({ value, onChange, onPressEnter, onClickButton }) => {

    return (
        <MyInput
            value={value}
            onChange={onChange}
            onPressEnter={onPressEnter}
            onClickButton={onClickButton}
            placeholder="Type a message..."
            icon="send"
            className='my-5'
            inputClassName='!p-3 rounded-lg !text-text !text-base'
            buttonClassName='!right-3 !p-1.5 my-button my-button-subtle !rounded-md !text-primary hover:!text-white'
        />
    )
}

export default ThreadMessageInput
