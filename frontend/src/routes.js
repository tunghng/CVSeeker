
// Not used yet
import React from "react";

export const MAIN_NAVIGATION = [
    {
        key: "home",
        path: "/",
        title: "CV Seeker",
        Component: React.lazy(() => import("./pages/HomePage")),
    },
    {
        key: "search",
        path: `/search/:id`,
        title: "Search | CV Seeker",
        Component: React.lazy(() => import("./pages/SearchPage")),
    },
    {
        key: "saved",
        path: `/saved`,
        title: "Saved | CV Seeker",
        Component: React.lazy(() => import("./pages/SavedPage")),
    },
    {
        key: "upload",
        path: `/upload`,
        title: "Upload CV | CV Seeker",
        Component: React.lazy(() => import("./pages/UploadPage")),
    },
];