class Robot {
public:
    int x, y;
    int _w, _h;
    int per, dir;
    vector<int> dx = {1, 0, -1, 0};
    vector<int> dy = {0, 1, 0, -1};
    vector<string> dirs = {"East", "North", "West", "South"};

    Robot(int width, int height) {
        x = 0;
        y = 0;
        dir = 0;
        _w = width;
        _h = height;
        per = 2 * (_w + _h) - 4;
    }
    
    void step(int num) {
        num = num % per;

        if(num == 0){
            num = per;
        }

        for(int step = 0; step < num; step++){
            int nx = x + dx[dir];
            int ny = y + dy[dir];

            if(nx < 0 || nx >= _w || ny < 0 || ny >= _h){
                dir = (dir + 1) % 4;
                nx = x + dx[dir];
                ny = y + dy[dir];
            }

            x = nx;
            y = ny;

        }
    }
    
    vector<int> getPos() {
        return {x, y};
    }
    
    string getDir() {
        return dirs[dir];
    }
};

/**
 * Your Robot object will be instantiated and called as such:
 * Robot* obj = new Robot(width, height);
 * obj->step(num);
 * vector<int> param_2 = obj->getPos();
 * string param_3 = obj->getDir();
 */