import type {Component} from 'solid-js';
import {Route, Routes} from "solid-app-router";
import Home from "./pages/Home";
import CreateBoard from "./pages/CreateBoard";
import PlayerSelection from "./pages/PlayerSelection";
import Game from "./pages/Game";

const App: Component = () => {
    return (
        <>
            <Routes>
                <Route path="/" element={Home}></Route>
                <Route path="/board/new" element={CreateBoard}></Route>
                <Route path="/player/selection" element={PlayerSelection}></Route>
                <Route path="/game" element={Game}></Route>
            </Routes>
        </>
    );
};

export default App;
