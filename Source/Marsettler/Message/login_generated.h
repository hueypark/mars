// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_LOGIN_FBS_H_
#define FLATBUFFERS_GENERATED_LOGIN_FBS_H_

#include "flatbuffers/flatbuffers.h"

namespace fbs {

struct Login;

FLATBUFFERS_MANUALLY_ALIGNED_STRUCT(8) Login FLATBUFFERS_FINAL_CLASS {
 private:
  int64_t Id_;

 public:
  Login() {
    memset(this, 0, sizeof(Login));
  }
  Login(int64_t _Id)
      : Id_(flatbuffers::EndianScalar(_Id)) {
  }
  int64_t Id() const {
    return flatbuffers::EndianScalar(Id_);
  }
};
FLATBUFFERS_STRUCT_END(Login, 8);

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_LOGIN_FBS_H_