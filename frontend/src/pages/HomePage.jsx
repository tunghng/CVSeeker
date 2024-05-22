
import { useState } from "react";
import { useNavigate } from "react-router-dom"

import ResumeSearchInput from "../components/ResumeSearchInput/ResumeSearchInput"
import ResumeSearchSlider from "../components/ResumeSearchSlider/ResumeSearchSlider"
import home_illustration from '../assets/images/home-illustration.png'

const HomePage = () => {
    // ====== State Management ======
    const navigate = useNavigate()
    const [resumeSearchInput, setResumeSearchInput] = useState('');
    const [resumeSearchLevel, setResumeSearchLevel] = useState(0.5);

    // ====== Event Handlers ======
    const resumeSearchKeyDownHandler = (e) => {
        if (e.key === 'Enter' && resumeSearchInput.trim() !== '') {
            navigate(`/search?query=${resumeSearchInput.trim()}&page=1&level=1`)
        }
    }
    const resumeSearchClickHandler = () => {
        if (resumeSearchInput.trim() !== '') {
            navigate(`/search?query=${resumeSearchInput.trim()}&page=1&level=1`)
        }
    }

    return (
        <main className="my-content-wrapper">
            {/* ====== Search Input ====== */}
            <div className="my-container-small pt-6">
                <ResumeSearchInput
                    value={resumeSearchInput}
                    onChange={(e) => setResumeSearchInput(e.target.value)}
                    onPressEnter={resumeSearchKeyDownHandler}
                    onClickButton={resumeSearchClickHandler}
                />
            </div>

            {/* ====== Search Slider ====== */}
            {/* <div className="my-container-small pt-3">
                <ResumeSearchSlider
                    value={resumeSearchLevel}
                    onChange={(e) => setResumeSearchLevel(e.target.value)}
                />
            </div> */}

            {/* ====== Illustration ====== */}
            {/* <div className="my-container-small mt-24 flex flex-col justify-center items-center">
                <img src={home_illustration} alt="Home Illustration" className="w-full max-w-96" />
            </div> */}
        </main>
    )
}

export default HomePage
