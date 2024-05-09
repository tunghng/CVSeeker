
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Layout from "./pages/Layout";
import HomePage from "./pages/HomePage";
import SearchPage from "./pages/SearchPage";
import SavedPage from "./pages/SavedPage";
import UploadPage from "./pages/UploadPage";

import { GlobalProvider } from "./contexts/GlobalContext";

import './styles/my-button.css';
import './styles/my-container.css';
import './styles/my-animation.css';

export default function App() {
    return (
        <GlobalProvider>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Layout />}>
                        <Route index element={<HomePage />} />
                        <Route path="search/:id" element={<SearchPage />} />
                        <Route path="upload" element={<UploadPage />} />
                        <Route path="saved" element={<SavedPage />} />
                        <Route path="*" element={<h1>Not Found</h1>} />
                    </Route>
                </Routes>
            </BrowserRouter>
        </GlobalProvider>
    )
}