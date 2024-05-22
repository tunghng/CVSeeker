
import MyInput from '../common/MyInput'

const ResumeSearchInput = ({ value, onChange, onPressEnter, onClickButton }) => {

    return (
        <MyInput
            value={value}
            onChange={onChange}
            onPressEnter={onPressEnter}
            onClickButton={onClickButton}
            placeholder="Search desired resume description..."
            icon="search"
        />
    );
};

export default ResumeSearchInput;
