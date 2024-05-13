
import MyInput from '../common/MyInput'

const ResumeSearchInput = ({ value, onChange, onPressEnter, onClickSearch }) => {

    return (
        <MyInput
            value={value}
            onChange={onChange}
            onPressEnter={onPressEnter}
            onClickSearch={onClickSearch}
            placeholder="Enter desired CV description..."
            icon="search"
        />
    );
};

export default ResumeSearchInput;
