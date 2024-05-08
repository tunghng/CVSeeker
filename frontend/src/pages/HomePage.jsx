
import SearchInput from "../components/SearchInput/SearchInput"
import SearchSlider from "../components/SearchSlider/SearchSlider"

const HomePage = () => {

    return (
        <main>
            {/* ====== Search Input ====== */}
            <div className="my-container-small pt-6">
                <SearchInput />
            </div>

            {/* ====== Search Slider ====== */}
            <div className="my-container-small pt-3">
                <SearchSlider />
            </div>
        </main>
    )
}

export default HomePage
